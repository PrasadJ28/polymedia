import type { SxProps } from '@mui/material/styles';

export const cardColors = {
  background: '#bababaff',
  darkShadow: 'rgba(163, 177, 198, 0.6)',
  lightShadow: 'rgba(255, 255, 255, 0.5)',
  text: '#002758ff',
};

export const cardStyle = (orientation: 'horizontal' | 'vertical', _height: number | 'auto'): SxProps => ({
  display: 'flex',
  flexDirection: orientation === 'horizontal' ? 'row' : 'column',
  alignItems: 'center',
  justifyContent: 'center',
  background: cardColors.background,
  borderRadius: '10px',
  border: 'none',
  boxShadow: 'none',
  transition: 'all 0.25s ease',
  cursor: 'pointer',

  '&:hover': {
    boxShadow: `6px 6px 12px ${cardColors.darkShadow}, -6px -6px 12px ${cardColors.lightShadow}`,
    transform: 'translateY(-3px)',
  },
  '&:active': {
    transform: 'translateY(2px)',
    boxShadow: `inset 6px 6px 12px ${cardColors.darkShadow}, inset -6px -6px 12px ${cardColors.lightShadow}`,
  },
});

export const textContainer: React.CSSProperties = {
  padding: '12px',
};

export const textClamp: SxProps = {
  color: cardColors.text,
  display: '-webkit-box',
  WebkitBoxOrient: 'vertical',
  WebkitLineClamp: 4,
  overflow: 'hidden',
  textOverflow: 'ellipsis',
};
