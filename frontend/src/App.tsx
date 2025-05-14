import { Suspense, lazy } from 'react';
import { BrowserRouter as Router, Routes, Route, Outlet } from 'react-router-dom';
import { Provider } from 'react-redux';
import CssBaseline from '@mui/material/CssBaseline';
import { store } from './store';
import { ColorModeProvider } from './providers/ColorModeProvider';

// Lazy load components
const Layout = lazy(() => import('./components/Layout'));
const Home = lazy(() => import('./pages/Home'));
const EventDetails = lazy(() => import('./pages/EventDetails'));
const Login = lazy(() => import('./pages/Login'));
const Register = lazy(() => import('./pages/Register'));
const Dashboard = lazy(() => import('./pages/admin/Dashboard'));
const TagsManagement = lazy(() => import('./pages/admin/TagsManagement'));
const ProtectedRoute = lazy(() => import('./components/ProtectedRoute'));

function App() {
  return (
    <Provider store={store}>
      <ColorModeProvider>
        <CssBaseline />
        <Router>
          <Suspense fallback={<div>Loading...</div>}>
            <Routes>
              <Route element={<Layout />}>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route path="/events/:id" element={<EventDetails />} />
                <Route
                  path="/admin"
                  element={<ProtectedRoute requiredRole="admin">
                    <Outlet />
                  </ProtectedRoute>}
                >
                  <Route path="dashboard" element={<Dashboard />} />
                  <Route path="tags" element={<TagsManagement />} />
                </Route>
              </Route>
            </Routes>
          </Suspense>
        </Router>
      </ColorModeProvider>
    </Provider>
  );
}

export default App;
