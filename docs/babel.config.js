module.exports = {
  plugins: [require.resolve("@babel/plugin-syntax-jsx")],
  presets: [
    require.resolve("@docusaurus/core/lib/babel/preset"),
    require.resolve("@babel/preset-react"),
  ],
};
