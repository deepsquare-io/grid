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
  themes: ['@docusaurus/theme-mermaid', 'docusaurus-theme-search-typesense'],

  presets: [
    [
      'classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: 'workflow',
          path: 'workflow',
          sidebarPath: require.resolve('./sidebars.js'),
          // Please change this to your repo.
          editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/docs/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],
  plugins: [
    [
      '@docusaurus/plugin-content-docs',
      {
        id: 'deepsquare-grid',
        path: 'deepsquare-grid',
        routeBasePath: 'deepsquare-grid',
        sidebarPath: require.resolve('./sidebars.js'),
        editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/docs/',
      },
    ],
    [
      '@docusaurus/plugin-content-docs',
      {
        id: 'blockchain',
        path: 'blockchain',
        routeBasePath: 'blockchain',
        sidebarPath: require.resolve('./sidebars.js'),
        editUrl: 'https://github.com/deepsquare-io/the-grid/tree/main/docs/',
      },
    ],
  ],
  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      typesense: {
        typesenseCollectionName: 'deepsquare-docs-index',
        typesenseServerConfig: {
          nodes: [
            {
              host: 'typesense.deepsquare.run',
              port: 443,
              protocol: 'https',
            },
          ],
          // Search only key. This is safe to share.
          apiKey: 'beEQxk6x8M8QyDHJpvcjpc56VWPF6JUi',
        },

        // Optional: Typesense search parameters: https://typesense.org/docs/0.24.0/api/search.html#search-parameters
        typesenseSearchParameters: {},

        // Optional
        contextualSearch: true,
      },
      tableOfContents: {
        maxHeadingLevel: 5,
      },
      docs: {
        sidebar: {
          autoCollapseCategories: false,
        },
      },
      navbar: {
        title: 'DeepSquare',
        logo: {
          alt: 'DeepSquare logo',
          src: 'img/logo.png',
        },
        items: [
          {
            href: '/workflow/overview',
            label: 'Workflow Documentation',
            position: 'left',
          },
          {
            href: '/deepsquare-grid/join/overview',
            label: 'Infrastructure Documentation',
            position: 'left',
          },
          {
            href: '/blockchain/introduction/overview',
            label: 'Blockchain Documentation',
            position: 'left',
          },
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
