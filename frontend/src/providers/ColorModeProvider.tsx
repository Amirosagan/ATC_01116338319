import { createContext, useContext, ReactNode } from 'react';
import { ThemeProvider } from '@mui/material/styles';
import { useColorMode } from '../hooks/useColorMode';

interface ColorModeContextType {
  mode: 'light' | 'dark';
  toggleColorMode: () => void;
}

export const ColorModeContext = createContext<ColorModeContextType>({
  mode: 'light',
  toggleColorMode: () => {},
});

interface ColorModeProviderProps {
  children: ReactNode;
}

export const ColorModeProvider = ({ children }: ColorModeProviderProps) => {
  const { theme, mode, toggleColorMode } = useColorMode();

  return (
    <ColorModeContext.Provider value={{ mode, toggleColorMode }}>
      <ThemeProvider theme={theme}>
        {children}
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
};

export const useColorModeContext = () => useContext(ColorModeContext); 