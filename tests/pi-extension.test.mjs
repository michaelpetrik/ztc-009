import assert from "node:assert/strict";
import { copyFile, mkdir, mkdtemp, readFile, realpath, rm, writeFile } from "node:fs/promises";
import { tmpdir } from "node:os";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import test from "node:test";
import register from "../extensions/ztc.js";

const root = fileURLToPath(new URL("../", import.meta.url));

function handler(registerExtension) {
  const handlers = new Map();
  registerExtension({
    on(name, callback) {
      assert.equal(name, "before_agent_start");
      assert.equal(handlers.has(name), false);
      handlers.set(name, callback);
    },
  });
  assert.equal(handlers.size, 1);
  return handlers.get("before_agent_start");
}

test("appends context without reflecting input or changing the event", () => {
  const invoke = handler(register);
  const event = Object.freeze({
    systemPrompt: "Original host and extension instructions.",
    prompt: "PRIVATE_INPUT_MARKER $(touch forbidden)",
    images: Object.freeze([]),
  });
  const result = invoke(event);
  assert.deepEqual(Object.keys(result), ["systemPrompt"]);
  assert.ok(result.systemPrompt.startsWith(event.systemPrompt + "\n\n"));
  assert.ok(result.systemPrompt.includes(JSON.stringify(path.join(root, "skills/ztc-response/SKILL.md"))));
  assert.equal(result.systemPrompt.includes("PRIVATE_INPUT_MARKER"), false);
  assert.deepEqual(invoke(event), result);
  assert.deepEqual(invoke({ ...event, prompt: "Unrelated user input" }), result);
  assert.equal(event.systemPrompt, "Original host and extension instructions.");
});

test("resolves bundled skill from a package path containing spaces", async (t) => {
  const temporary = await mkdtemp(path.join(tmpdir(), "ztc pi "));
  t.after(() => rm(temporary, { recursive: true, force: true }));
  const packageRoot = path.join(temporary, "package with spaces");
  await mkdir(path.join(packageRoot, "extensions"), { recursive: true });
  await mkdir(path.join(packageRoot, "skills/ztc-response"), { recursive: true });
  await writeFile(path.join(packageRoot, "package.json"), '{"type":"module"}');
  const skill = path.join(packageRoot, "skills/ztc-response/SKILL.md");
  await writeFile(skill, "Test skill");
  const extension = path.join(packageRoot, "extensions/ztc.js");
  await copyFile(path.join(root, "extensions/ztc.js"), extension);
  const loaded = await import(pathToFileURL(extension).href);
  const result = handler(loaded.default)({ systemPrompt: "Host" });
  assert.ok(result.systemPrompt.includes(JSON.stringify(await realpath(skill))));
  assert.equal(await readFile(skill, "utf8"), "Test skill");
});

test("manifest exposes existing skills and the tested extension without dependencies", async () => {
  const manifest = JSON.parse(await readFile(path.join(root, "package.json"), "utf8"));
  assert.equal(manifest.type, "module");
  assert.ok(manifest.keywords.includes("pi-package"));
  assert.deepEqual(manifest.pi.extensions, ["./extensions/ztc.js"]);
  assert.deepEqual(manifest.pi.skills, ["./skills"]);
  assert.equal(manifest.dependencies, undefined);
  for (const name of ["ztc-install", "ztc-response"]) {
    assert.ok((await readFile(path.join(root, "skills", name, "SKILL.md"), "utf8")).startsWith("---\n"));
  }
});
