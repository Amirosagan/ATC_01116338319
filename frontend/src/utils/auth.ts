const TOKEN_KEY = 'event_booking_token';

export const getToken = (): string | null => {
  return localStorage.getItem(TOKEN_KEY);
};

export const setToken = (token: string): void => {
  localStorage.setItem(TOKEN_KEY, token);
};

export const removeToken = (): void => {
  localStorage.removeItem(TOKEN_KEY);
};

export const getAuthHeader = (): { Authorization: string } | undefined => {
  const token = getToken();
  return token ? { Authorization: `Bearer ${token}` } : undefined;
}; 