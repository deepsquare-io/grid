import { CssVarsProvider, extendTheme, getInitColorSchemeScript } from '@mui/joy';
import React from 'react';

const theme = extendTheme({
  fontFamily: {
    body: 'var(--custom-font-base)',
    code: 'var(--custom-font-code)',
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
