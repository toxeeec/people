module.exports = {
  extends: [
    'airbnb-typescript/base',
    'prettier',
    'prettier/@typescript-eslint',
    'plugin:prettier/recommended',
  ],
  ignorePatterns: ['dist', 'jest.config.js'],
  parserOptions: {
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
  },
  rules: {
    'prettier/prettier': [
      'warn',
      {
        endOfLine: 'auto',
        singleQuote: true,
        trailingComma: 'es5',
      },
    ],
    'import/prefer-default-export': 'off',
    'no-console': 'off',
    'consistent-return': 'off',
    '@typescript-eslint/no-unused-vars': ['error', { args: 'none' }],
  },
};
