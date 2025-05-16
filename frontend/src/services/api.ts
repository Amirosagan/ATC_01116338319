import type { Event, User, Tag } from '../types';
import axiosInstance from './axiosInstance';

interface LoginCredentials {
  email: string;
  password: string;
}

interface RegisterCredentials extends LoginCredentials {
  username: string;
}

interface AuthResponse {
  user: User;
  token: string;
}

interface ApiError {
  error: string;
}

interface CreateEventData {
  name: string;
  description: string;
  category: string;
  date: string;
  location: string;
  price: number;
  image: string;
  tagIds: string[];
}

interface CreateTagData {
  name: string;
}

interface BookingResponse {
  id: string;
  userId: string;
  eventId: string;
  event: Event;
  createdAt: string;
}

interface UploadResponse {
  imageUrl: string;
}

const handleApiError = (error: any): never => {
  if (error.response?.data?.error) {
    throw new Error(error.response.data.error);
  }
  throw new Error('An unexpected error occurred');
};

const api = {
  auth: {
    login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
      try {
        const { data } = await axiosInstance.post<AuthResponse>('/auth/login', credentials);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    register: async (credentials: RegisterCredentials): Promise<AuthResponse> => {
      try {
        const { data } = await axiosInstance.post<AuthResponse>('/auth/register', credentials);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },
  },

  events: {
    getAll: async (): Promise<Event[]> => {
      try {
        const { data } = await axiosInstance.get<Event[]>('/events');
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    getById: async (id: string): Promise<Event> => {
      try {
        const { data } = await axiosInstance.get<Event>(`/events/${id}`);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    create: async (eventData: CreateEventData): Promise<Event> => {
      try {
        const { data } = await axiosInstance.post<Event>('/events', eventData);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    update: async (id: string, eventData: Partial<CreateEventData>): Promise<Event> => {
      try {
        const { data } = await axiosInstance.put<Event>(`/events/${id}`, eventData);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    delete: async (id: string): Promise<void> => {
      try {
        await axiosInstance.delete(`/events/${id}`);
      } catch (error) {
        throw handleApiError(error);
      }
    },
  },

  bookings: {
    create: async (eventId: string): Promise<BookingResponse> => {
      try {
        const { data } = await axiosInstance.post<BookingResponse>('/bookings', { eventId });
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    getUserBookings: async (): Promise<BookingResponse[]> => {
      try {
        const { data } = await axiosInstance.get<BookingResponse[]>('/bookings/user');
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },
  },

  tags: {
    getAll: async (): Promise<Tag[]> => {
      try {
        const { data } = await axiosInstance.get<Tag[]>('/tags');
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    getById: async (id: string): Promise<Tag> => {
      try {
        const { data } = await axiosInstance.get<Tag>(`/tags/${id}`);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    create: async (tagData: CreateTagData): Promise<Tag> => {
      try {
        const { data } = await axiosInstance.post<Tag>('/tags', tagData);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    update: async (id: string, tagData: CreateTagData): Promise<Tag> => {
      try {
        const { data } = await axiosInstance.put<Tag>(`/tags/${id}`, tagData);
        return data;
      } catch (error) {
        throw handleApiError(error);
      }
    },

    delete: async (id: string): Promise<void> => {
      try {
        await axiosInstance.delete(`/tags/${id}`);
      } catch (error) {
        throw handleApiError(error);
      }
    },
  },

  upload: {
    image: async (file: File): Promise<string> => {
      try {
        const formData = new FormData();
        formData.append('image', file);

        const { data } = await axiosInstance.post<UploadResponse>('/upload/image', formData, {
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        });

        const apiHost = "https://areeb.ddns.net:2025";

        return apiHost + data.imageUrl;
      } catch (error) {
        throw handleApiError(error);
      }
    },
  },
};

export default api; 