import { useState, useEffect } from 'react';
import {
  Box,
  Button,
  Container,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  OutlinedInput,
  Alert,
  IconButton,
} from '@mui/material';
import { PhotoCamera } from '@mui/icons-material';
import { useFormik } from 'formik';
import * as yup from 'yup';
import type { Event, Tag } from '../../types';
import api from '../../services/api';

const MAX_FILE_SIZE = 3 * 1024 * 1024; // 3MB
const ALLOWED_FILE_TYPES = ['image/jpeg', 'image/png', 'image/gif'];

const validationSchema = yup.object({
  name: yup.string().required('Name is required'),
  description: yup.string().required('Description is required'),
  category: yup.string().required('Category is required'),
  date: yup.string().required('Date is required'),
  location: yup.string().required('Location is required'),
  price: yup.number().required('Price is required').min(0, 'Price must be positive'),
  image: yup.string().required('Image URL is required'),
  tagIds: yup.array().of(yup.string()).min(1, 'At least one tag is required'),
});

const Dashboard = () => {
  const [events, setEvents] = useState<Event[]>([]);
  const [tags, setTags] = useState<Tag[]>([]);
  const [open, setOpen] = useState(false);
  const [editingEvent, setEditingEvent] = useState<Event | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [uploadProgress, setUploadProgress] = useState<boolean>(false);

  useEffect(() => {
    fetchEvents();
    fetchTags();
  }, []);

  const fetchEvents = async () => {
    try {
      const data = await api.events.getAll();
      setEvents(data);
    } catch (error) {
      setError(error instanceof Error ? error.message : 'Failed to fetch events');
    }
  };

  const fetchTags = async () => {
    try {
      const data = await api.tags.getAll();
      setTags(data);
    } catch (error) {
      setError(error instanceof Error ? error.message : 'Failed to fetch tags');
    }
  };

  const handleClose = () => {
    setOpen(false);
    setEditingEvent(null);
    setError(null);
    formik.resetForm();
  };

  const formatDateToUTC = (date: string) => {
    const d = new Date(date);
    return d.toISOString();
  };

  const handleImageUpload = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (!file) return;

    // Validate file size
    if (file.size > MAX_FILE_SIZE) {
      setError('File size exceeds 3MB limit');
      return;
    }

    // Validate file type
    if (!ALLOWED_FILE_TYPES.includes(file.type)) {
      setError('Invalid file type. Please upload a JPG, PNG, or GIF image.');
      return;
    }

    try {
      setUploadProgress(true);
      const imageUrl = await api.upload.image(file);
      formik.setFieldValue('image', imageUrl);
      setError(null);
    } catch (error) {
      setError(error instanceof Error ? error.message : 'Failed to upload image');
    } finally {
      setUploadProgress(false);
    }
  };

  const formik = useFormik({
    initialValues: {
      name: '',
      description: '',
      category: '',
      date: '',
      location: '',
      price: '',
      image: '',
      tagIds: [] as string[],
    },
    validationSchema: validationSchema,
    onSubmit: async (values) => {
      try {
        const formattedData = {
          ...values,
          price: Number(values.price),
          date: formatDateToUTC(values.date),
          tagIds: values.tagIds,
        };

        if (editingEvent) {
          await api.events.update(editingEvent.id, formattedData);
        } else {
          await api.events.create(formattedData);
        }
        handleClose();
        fetchEvents();
      } catch (error) {
        setError(error instanceof Error ? error.message : 'Failed to save event');
      }
    },
  });

  const handleEdit = (event: Event) => {
    setEditingEvent(event);
    formik.setValues({
      name: event.name,
      description: event.description,
      category: event.category,
      date: new Date(event.date).toISOString().slice(0, 16),
      location: event.location,
      price: event.price.toString(),
      image: event.image,
      tagIds: event.tags?.map(tag => tag.id) || [],
    });
    setOpen(true);
  };

  const handleDelete = async (id: string) => {
    if (window.confirm('Are you sure you want to delete this event?')) {
      try {
        await api.events.delete(id);
        fetchEvents();
      } catch (error) {
        setError(error instanceof Error ? error.message : 'Failed to delete event');
      }
    }
  };

  return (
    <Container>
      <Box sx={{ mt: 4, mb: 4 }}>
        {error && (
          <Alert severity="error" sx={{ mb: 2 }} onClose={() => setError(null)}>
            {error}
          </Alert>
        )}

        <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
          <Typography variant="h4" component="h1">
            Event Management
          </Typography>
          <Button variant="contained" onClick={() => setOpen(true)}>
            Add New Event
          </Button>
        </Box>

        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Name</TableCell>
                <TableCell>Category</TableCell>
                <TableCell>Date</TableCell>
                <TableCell>Location</TableCell>
                <TableCell>Price</TableCell>
                <TableCell>Tags</TableCell>
                <TableCell>Image</TableCell>
                <TableCell>Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {events.map((event) => (
                <TableRow key={event.id}>
                  <TableCell>{event.name}</TableCell>
                  <TableCell>{event.category}</TableCell>
                  <TableCell>
                    {new Date(event.date).toLocaleString()}
                  </TableCell>
                  <TableCell>{event.location}</TableCell>
                  <TableCell>${event.price}</TableCell>
                  <TableCell>
                    {event.tags?.map(tag => tag.name).join(', ')}
                  </TableCell>
                  <TableCell>
                    <img 
                      src={event.image} 
                      alt={event.name} 
                      style={{ width: '50px', height: '50px', objectFit: 'cover' }}
                    />
                  </TableCell>
                  <TableCell>
                    <Button
                      size="small"
                      onClick={() => handleEdit(event)}
                      sx={{ mr: 1 }}
                    >
                      Edit
                    </Button>
                    <Button
                      size="small"
                      color="error"
                      onClick={() => handleDelete(event.id)}
                    >
                      Delete
                    </Button>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>

        <Dialog open={open} onClose={handleClose} maxWidth="sm" fullWidth>
          <DialogTitle>
            {editingEvent ? 'Edit Event' : 'Add New Event'}
          </DialogTitle>
          <form onSubmit={formik.handleSubmit}>
            <DialogContent>
              <TextField
                fullWidth
                margin="normal"
                name="name"
                label="Event Name"
                value={formik.values.name}
                onChange={formik.handleChange}
                error={formik.touched.name && Boolean(formik.errors.name)}
                helperText={formik.touched.name && formik.errors.name}
              />
              <TextField
                fullWidth
                margin="normal"
                name="description"
                label="Description"
                multiline
                rows={4}
                value={formik.values.description}
                onChange={formik.handleChange}
                error={formik.touched.description && Boolean(formik.errors.description)}
                helperText={formik.touched.description && formik.errors.description}
              />
              <TextField
                fullWidth
                margin="normal"
                name="category"
                label="Category"
                value={formik.values.category}
                onChange={formik.handleChange}
                error={formik.touched.category && Boolean(formik.errors.category)}
                helperText={formik.touched.category && formik.errors.category}
              />
              <TextField
                fullWidth
                margin="normal"
                name="date"
                label="Date"
                type="datetime-local"
                value={formik.values.date}
                onChange={formik.handleChange}
                error={formik.touched.date && Boolean(formik.errors.date)}
                helperText={formik.touched.date && formik.errors.date}
                InputLabelProps={{
                  shrink: true,
                }}
              />
              <TextField
                fullWidth
                margin="normal"
                name="location"
                label="Location"
                value={formik.values.location}
                onChange={formik.handleChange}
                error={formik.touched.location && Boolean(formik.errors.location)}
                helperText={formik.touched.location && formik.errors.location}
              />
              <TextField
                fullWidth
                margin="normal"
                name="price"
                label="Price"
                type="number"
                value={formik.values.price}
                onChange={formik.handleChange}
                error={formik.touched.price && Boolean(formik.errors.price)}
                helperText={formik.touched.price && formik.errors.price}
              />
              <Box sx={{ mt: 2, mb: 2 }}>
                <input
                  accept="image/*"
                  style={{ display: 'none' }}
                  id="image-upload"
                  type="file"
                  onChange={handleImageUpload}
                />
                <label htmlFor="image-upload">
                  <IconButton
                    color="primary"
                    component="span"
                    disabled={uploadProgress}
                  >
                    <PhotoCamera />
                  </IconButton>
                  {uploadProgress ? 'Uploading...' : 'Upload Image'}
                </label>
                {formik.values.image && (
                  <Box sx={{ mt: 1 }}>
                    <img
                      src={formik.values.image}
                      alt="Preview"
                      style={{ width: '100px', height: '100px', objectFit: 'cover' }}
                    />
                  </Box>
                )}
              </Box>
              <FormControl 
                fullWidth 
                margin="normal" 
                error={formik.touched.tagIds && Boolean(formik.errors.tagIds)}
              >
                <InputLabel id="tags-label">Tags *</InputLabel>
                <Select
                  labelId="tags-label"
                  multiple
                  name="tagIds"
                  value={formik.values.tagIds}
                  onChange={formik.handleChange}
                  input={<OutlinedInput label="Tags *" />}
                  renderValue={(selected) => (
                    <Box sx={{ display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                      {selected.map((value) => {
                        const tag = tags.find(t => t.id === value);
                        return (
                          <Box
                            key={value}
                            sx={{
                              bgcolor: 'primary.main',
                              color: 'white',
                              px: 1,
                              py: 0.5,
                              borderRadius: 1,
                              fontSize: '0.8rem'
                            }}
                          >
                            {tag?.name}
                          </Box>
                        );
                      })}
                    </Box>
                  )}
                >
                  {tags.map((tag) => (
                    <MenuItem key={tag.id} value={tag.id}>
                      {tag.name}
                    </MenuItem>
                  ))}
                </Select>
                {formik.touched.tagIds && formik.errors.tagIds && (
                  <Typography color="error" variant="caption" sx={{ mt: 0.5 }}>
                    {formik.errors.tagIds as string}
                  </Typography>
                )}
              </FormControl>
            </DialogContent>
            <DialogActions>
              <Button onClick={handleClose}>Cancel</Button>
              <Button 
                type="submit" 
                variant="contained"
                disabled={uploadProgress}
              >
                {editingEvent ? 'Update' : 'Create'}
              </Button>
            </DialogActions>
          </form>
        </Dialog>
      </Box>
    </Container>
  );
};

export default Dashboard; 