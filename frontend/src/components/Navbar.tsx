import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import {
  AppBar,
  Box,
  Toolbar,
  IconButton,
  Typography,
  Menu,
  MenuItem,
  Button,
  useTheme,
  Avatar,
  Tooltip,
} from '@mui/material';
import {
  Menu as MenuIcon,
  AdminPanelSettings,
} from '@mui/icons-material';
import useAuth from '../hooks/useAuth';

interface NavbarProps {
  children?: React.ReactNode;
}

const Navbar = ({ children }: NavbarProps) => {
  const navigate = useNavigate();
  const theme = useTheme();
  const { isAuthenticated, user, logout } = useAuth();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [mobileMenuAnchor, setMobileMenuAnchor] = useState<null | HTMLElement>(null);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleMobileMenu = (event: React.MouseEvent<HTMLElement>) => {
    setMobileMenuAnchor(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
    setMobileMenuAnchor(null);
  };

  const handleNavigate = (path: string) => {
    navigate(path);
    handleClose();
  };

  return (
    <AppBar 
      position="fixed" 
      elevation={0}
      sx={{
        backdropFilter: 'blur(6px)',
        backgroundColor: 'background.paper',
        borderBottom: `1px solid ${theme.palette.divider}`,
      }}
    >
      <Toolbar>
        <IconButton
          size="large"
          edge="start"
          color="inherit"
          aria-label="menu"
          sx={{ mr: 2, display: { sm: 'none' } }}
          onClick={handleMobileMenu}
        >
          <MenuIcon />
        </IconButton>

        <Typography
          variant="h6"
          component="div"
          sx={{
            flexGrow: 1,
            cursor: 'pointer',
            fontWeight: 600,
            background: theme.palette.mode === 'dark'
              ? 'linear-gradient(45deg, #90caf9 30%, #ce93d8 90%)'
              : 'linear-gradient(45deg, #1976d2 30%, #9c27b0 90%)',
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
          }}
          onClick={() => navigate('/')}
        >
          Event Booking
        </Typography>

        <Box sx={{ display: { xs: 'none', sm: 'flex' }, alignItems: 'center', gap: 2 }}>
          {children}
          {isAuthenticated ? (
            <>
              {user?.role === 'admin' && (
                <Tooltip title="Admin Panel">
                  <IconButton
                    color="inherit"
                    onClick={handleMenu}
                    sx={{ ml: 1 }}
                  >
                    <AdminPanelSettings />
                  </IconButton>
                </Tooltip>
              )}
              <Tooltip title="Account">
                <IconButton onClick={handleMenu} sx={{ p: 0 }}>
                  <Avatar sx={{ bgcolor: theme.palette.primary.main }}>
                    {user?.username?.[0]?.toUpperCase() || 'U'}
                  </Avatar>
                </IconButton>
              </Tooltip>
            </>
          ) : (
            <>
              <Button
                variant="outlined"
                color="inherit"
                onClick={() => navigate('/login')}
              >
                Login
              </Button>
              <Button
                variant="contained"
                color="primary"
                onClick={() => navigate('/register')}
              >
                Register
              </Button>
            </>
          )}
        </Box>

        {/* Admin Menu */}
        <Menu
          anchorEl={anchorEl}
          open={Boolean(anchorEl)}
          onClose={handleClose}
          onClick={handleClose}
          PaperProps={{
            sx: {
              mt: 1.5,
              overflow: 'visible',
              filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
              '&:before': {
                content: '""',
                display: 'block',
                position: 'absolute',
                top: 0,
                right: 14,
                width: 10,
                height: 10,
                bgcolor: 'background.paper',
                transform: 'translateY(-50%) rotate(45deg)',
                zIndex: 0,
              },
            },
          }}
          transformOrigin={{ horizontal: 'right', vertical: 'top' }}
          anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
        >
          {user?.role === 'admin' && (
            <>
              <MenuItem onClick={() => handleNavigate('/admin/dashboard')}>
                Event Management
              </MenuItem>
              <MenuItem onClick={() => handleNavigate('/admin/tags')}>
                Tag Management
              </MenuItem>
            </>
          )}
          <MenuItem onClick={logout}>Logout</MenuItem>
        </Menu>

        {/* Mobile Menu */}
        <Menu
          anchorEl={mobileMenuAnchor}
          open={Boolean(mobileMenuAnchor)}
          onClose={handleClose}
          PaperProps={{
            sx: {
              width: '100%',
              maxWidth: '300px',
              mt: 1.5,
            },
          }}
        >
          {isAuthenticated ? (
            <>
              {user?.role === 'admin' && (
                <>
                  <MenuItem onClick={() => handleNavigate('/admin/dashboard')}>
                    Event Management
                  </MenuItem>
                  <MenuItem onClick={() => handleNavigate('/admin/tags')}>
                    Tag Management
                  </MenuItem>
                </>
              )}
              <MenuItem onClick={logout}>Logout</MenuItem>
            </>
          ) : (
            <>
              <MenuItem onClick={() => handleNavigate('/login')}>
                Login
              </MenuItem>
              <MenuItem onClick={() => handleNavigate('/register')}>
                Register
              </MenuItem>
            </>
          )}
        </Menu>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar; 