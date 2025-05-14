import { Alert, Container } from '@mui/material';

interface ErrorMessageProps {
  message: string;
}

const ErrorMessage = ({ message }: ErrorMessageProps) => {
  return (
    <Container>
      <Alert severity="error" sx={{ mt: 4 }}>
        {message}
      </Alert>
    </Container>
  );
};

export default ErrorMessage; 