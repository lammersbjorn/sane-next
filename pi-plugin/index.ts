import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";
import type { ExtensionAPI } from "@mariozechner/pi-coding-agent";

const baseDir = dirname(fileURLToPath(import.meta.url));

export default function (pi: ExtensionAPI) {
  pi.on("resources_discover", () => ({
    skillPaths: [join(baseDir, "..", "packs")],
  }));

  pi.registerCommand("sane-status", {
    description: "Show Sane overlay status and discovered pack location",
    handler: async (_args, ctx) => {
      ctx.ui.notify(`Sane overlay loaded from ${baseDir}`, "info");
    },
  });
}
