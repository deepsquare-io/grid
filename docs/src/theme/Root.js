import { CssVarsProvider, extendTheme, getInitColorSchemeScript } from '@mui/joy';
import React from 'react';

const theme = extendTheme({
  fontFamily: {
    body: 'var(--custom-font-base)',
    code: 'var(--custom-font-code)',
  },
  colorSchemes: {
    light: {
      palette: {
        primary: {
          50: '#dea1fe',
          100: '#d78efe',
          200: '#d17bfe',
          300: '#ca69fd',
          400: '#c456fd',
          500: '#bd43fd',
          600: '#aa3ce4',
          700: '#9736ca',
          800: '#842fb1',
          900: '#712898',
        },
      },
    },
    dark: {
      palette: {
        primary: {
          50: '#dea1fe',
          100: '#d78efe',
          200: '#d17bfe',
          300: '#ca69fd',
          400: '#c456fd',
          500: '#bd43fd',
          600: '#aa3ce4',
          700: '#9736ca',
          800: '#842fb1',
          900: '#712898',
        },
      },
    },
  },
});

// Default implementation, that you can customize
export default function Root({ children }) {
  return (
    <>
      {getInitColorSchemeScript()}
      <CssVarsProvider theme={theme}>{children}</CssVarsProvider>
    </>
  );
}
