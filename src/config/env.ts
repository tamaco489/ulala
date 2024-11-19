export const API_HOST =
  process.env['NEXT_PUBLIC_API_HOST'] ?? 'http://localhost:8080';
export const API_REQUEST_OPTIONS = {
  headers: {
    Authorization: `Bearer ${
      process.env['NEXT_PUBLIC_API_KEY'] ?? 'ABCDEFG123456789'
    }`,
    'Content-Type': 'application/json',
  },
  withCredentials: true,
};
