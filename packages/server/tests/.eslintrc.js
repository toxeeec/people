module.exports = {
  parserOptions: {
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
  },
  extends: ['plugin:jest/style'],
  env: {
    'jest/globals': true,
  },
  rules: {
    'no-restricted-syntax': 'off',
  },
};
