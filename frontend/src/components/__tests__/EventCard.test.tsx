import { render, screen, fireEvent } from '@testing-library/react';
import { describe, it, expect, vi } from 'vitest';
import EventCard from '../EventCard';

describe('EventCard', () => {
  const mockEvent = {
    id: 1,
    title: 'Test Event',
    description: 'Test Description',
    startTime: '2024-03-20T10:00:00Z',
    endTime: '2024-03-20T12:00:00Z',
    location: 'Test Location',
    creatorId: 1
  };

  const mockOnClick = vi.fn();

  it('renders event details correctly', () => {
    render(<EventCard event={mockEvent} onClick={mockOnClick} />);

    expect(screen.getByText(mockEvent.title)).toBeInTheDocument();
    expect(screen.getByText(mockEvent.description)).toBeInTheDocument();
    expect(screen.getByText(mockEvent.location)).toBeInTheDocument();
  });

  it('calls onClick handler when clicked', () => {
    render(<EventCard event={mockEvent} onClick={mockOnClick} />);

    const card = screen.getByRole('article');
    fireEvent.click(card);

    expect(mockOnClick).toHaveBeenCalledWith(mockEvent.id);
  });

  it('displays formatted date and time', () => {
    render(<EventCard event={mockEvent} onClick={mockOnClick} />);

    const dateTimeString = screen.getByText(/March 20, 2024/);
    expect(dateTimeString).toBeInTheDocument();
  });
}); 