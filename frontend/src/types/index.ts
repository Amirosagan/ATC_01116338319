export interface User {
  id: string;
  username: string;
  email: string;
  role: 'admin' | 'user';
}

export interface Event {
  id: string;
  name: string;
  description: string;
  category: string;
  date: string;
  location: string;
  price: number;
  image: string;
  tags?: Tag[];
}

export interface Booking {
  id: string;
  userId: string;
  eventId: string;
  bookingDate: string;
}

export interface Tag {
  id: string;
  name: string;
}

export interface AuthState {
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
} 