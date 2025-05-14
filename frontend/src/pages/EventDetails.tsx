import { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import {
  Box,
  Button,
  Card,
  CardMedia,
  Container,
  Grid,
  Typography,
  Chip,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
} from '@mui/material';
import { useSelector } from 'react-redux';
import type { RootState } from '../store';
import type { Event } from '../types';
import api from '../services/api';

const EventDetails = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [event, setEvent] = useState<Event | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [showSuccess, setShowSuccess] = useState(false);
  const { isAuthenticated } = useSelector((state: RootState) => state.auth);

  useEffect(() => {
    const fetchEvent = async () => {
      try {
        const data = await api.events.getById(id!);
        setEvent(data);
        setError(null);
      } catch (error) {
        console.error('Error fetching event:', error);
        setError('Failed to load event details');
      } finally {
        setLoading(false);
      }
    };

    fetchEvent();
  }, [id]);

  const handleBooking = async () => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }

    try {
      await api.bookings.create(id!);
      setShowSuccess(true);
      setError(null);
    } catch (error) {
      console.error('Error booking event:', error);
      setError(error instanceof Error ? error.message : 'Failed to book event');
    }
  };

  if (loading) {
    return (
      <Box sx={{ textAlign: 'center', py: 4 }}>
        <Typography>Loading event details...</Typography>
      </Box>
    );
  }

  if (error) {
    return (
      <Box sx={{ textAlign: 'center', py: 4 }}>
        <Typography color="error">{error}</Typography>
      </Box>
    );
  }

  if (!event) {
    return (
      <Box sx={{ textAlign: 'center', py: 4 }}>
        <Typography>Event not found</Typography>
      </Box>
    );
  }

  return (
    <Container maxWidth="lg">
      <Grid container spacing={4}>
        <Grid
          item
          xs={12}
          md={6}
          // @ts-ignore - MUI Grid typing issue
        >
          <Card>
            <CardMedia
              component="img"
              height="400"
              image={event.image}
              alt={event.name}
              sx={{ objectFit: 'cover' }}
            />
          </Card>
        </Grid>
        <Grid
          item
          xs={12}
          md={6}
          // @ts-ignore - MUI Grid typing issue
        >
          <Typography variant="h4" component="h1" gutterBottom>
            {event.name}
          </Typography>
          <Typography variant="h6" color="text.secondary" gutterBottom>
            {new Date(event.date).toLocaleDateString()} at {event.location}
          </Typography>
          <Box sx={{ my: 2 }}>
            {event.tags?.map((tag) => (
              <Chip
                key={tag.id}
                label={tag.name}
                sx={{ mr: 1, mb: 1 }}
                variant="outlined"
              />
            ))}
          </Box>
          <Typography variant="body1" paragraph>
            {event.description}
          </Typography>
          <Typography variant="h6" gutterBottom>
            Price: ${event.price}
          </Typography>
          <Button
            variant="contained"
            size="large"
            fullWidth
            onClick={handleBooking}
            sx={{ mt: 2 }}
          >
            Book Now
          </Button>
        </Grid>
      </Grid>

      <Dialog open={showSuccess} onClose={() => setShowSuccess(false)}>
        <DialogTitle>Booking Successful!</DialogTitle>
        <DialogContent>
          <Typography>
            You have successfully booked a ticket for {event.name}. We look forward
            to seeing you at the event!
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => navigate('/')} color="primary">
            Back to Events
          </Button>
        </DialogActions>
      </Dialog>
    </Container>
  );
};

export default EventDetails; 