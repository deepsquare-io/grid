// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require('prism-react-renderer').themes.github;
const darkCodeTheme = require('prism-react-renderer').themes.palenight;

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'DeepSquare',
  tagline: 'The decentralized High Performance Computing Ecosystem',
  url: 'https://docs.deepsquare.run',
  baseUrl: '/',
  onBrokenLinks: 'ignore',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/logo.png',
  organizationName: 'DeepSquare',
  projectName: 'DeepSquare',
  trailingSlash: false,
  markdown: {
    mermaid: true,
  },
  themes: ['@docusaurus/theme-mermaid'],

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/docs/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],
  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      tableOfContents: {
        maxHeadingLevel: 5,
      },
      docs: {
        sidebar: {
          autoCollapseCategories: false,
        },
      },
      navbar: {
        title: 'DeepSquare Documentation',
        logo: {
          alt: 'DeepSquare logo',
          src: 'img/logo.png',
        },
        items: [
          {
            href: 'https://discord.gg/zvFnqVHmJh',
            label: 'Discord',
            position: 'right',
          },
          {
            href: 'https://github.com/deepsquare-io/the-grid',
            label: 'GitHub',
            position: 'right',
          },
        ],
      },
      footer: {
        copyright: `<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">CC BY-SA 4.0</a> © ${new Date().getFullYear()} <a href="https://deepsquare.io">DeepSquare</a> documentation. Built with Docusaurus.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
        additionalLanguages: ['properties', 'docker'],
      },
      colorMode: {
        defaultMode: 'light',
        disableSwitch: false,
        respectPrefersColorScheme: false,
      },
    }),
};

module.exports = config;
