import { navbar } from "vuepress-theme-hope";

export default navbar([
  "/",
  {
    text: "Documentation",
    icon: "book",
    children: [
      {
        text: "Shortcuts",
        prefix: "/",
        children: [
          { text: "Install", link: "install.html", icon: "arrow-down" },
          { text: "Language Tour", link: "tour.html", icon: "life-ring" },
          { text: "Rosetta Stone", link: "user-guide/rosetta-stone.html", icon: "language" },
          { text: "User Guide", link: "user-guide/", icon: "book" },
          { text: "Integrations", link: "integrations/", icon: "puzzle-piece" },
          { text: "Operators And Tokens", link: "parser/", icon: "equals" },
          { text: "Builtins", link: "commands/", icon: "terminal" },
          { text: "Variables", link: "variables/", icon: "dollar" },
          { text: "Data Types", link: "types/", icon: "table" },
          { text: "Events", link: "events/", icon: "bolt" },
        ],
      },
    ],
  },
  "/changelog/",
  "/blog/",
  { text: "Discuss", link: "https://github.com/lmorg/murex/discussions", icon: "comment" },
  "/contributing",
]);
