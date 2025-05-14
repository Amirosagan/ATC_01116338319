import {
  Box,
  Container,
  useTheme,
  useMediaQuery,
  IconButton,
} from '@mui/material';
import { Brightness4, Brightness7 } from '@mui/icons-material';
import { Outlet } from 'react-router-dom';
import Navbar from './Navbar';
import { useColorModeContext } from '../providers/ColorModeProvider';

const Layout = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));
  const { mode, toggleColorMode } = useColorModeContext();

  return (
    <Box
      sx={{
        minHeight: '100vh',
        bgcolor: 'background.default',
        color: 'text.primary',
        transition: 'all 0.3s',
      }}
    >
      <Navbar>
        <IconButton
          sx={{ ml: 1 }}
          onClick={toggleColorMode}
          color="inherit"
          aria-label="toggle dark mode"
        >
          {mode === 'dark' ? <Brightness7 /> : <Brightness4 />}
        </IconButton>
      </Navbar>
      
      <Container
        maxWidth="lg"
        sx={{
          py: isMobile ? 2 : 4,
          mt: 8, // Space for fixed navbar
        }}
      >
        <Outlet />
      </Container>
    </Box>
  );
};

export default Layout; 