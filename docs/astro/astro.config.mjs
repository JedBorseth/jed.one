// @ts-check
import { defineConfig, passthroughImageService } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
  image: {
    service: passthroughImageService(),
  },
  outDir: "../static",
  integrations: [
    starlight({
      title: "Jed.One",
      social: [
        {
          icon: "github",
          label: "GitHub",
          href: "https://github.com/jedborseth",
        },
        {
          icon: "twitter",
          label: "Twitter",
          href: "https://twitter.com/jedborseth",
        },
      ],
      sidebar: [
        {
          label: "Guides",
          items: [
            // Each item here is one entry in the navigation menu.
            { label: "Getting Started", slug: "guides/example" },
          ],
        },
        {
          label: "Routes",
          autogenerate: { directory: "routes" },
        },
      ],
    }),
  ],
});
