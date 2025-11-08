import type { SxProps } from '@mui/material/styles';

export const barColors = {
  background: '#E6E9EF',
  darkShadow: 'rgba(163, 177, 198, 0.6)',
  lightShadow: 'rgba(255, 255, 255, 0.5)',
  text: '#002758',
};


export const barContainer: SxProps = {
  width: '100%',
  height: 70,
  background: barColors.background,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'space-between',
  padding: '0 2rem',
  boxShadow: 'none',
  border: 'none',
  position: 'sticky',
  top: 0,
  zIndex: 100,
};

export const elevatedCircle: SxProps = {
  width: 50,
  height: 50,
  borderRadius: '50%',
  background: barColors.background,
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  boxShadow: `6px 6px 12px ${barColors.darkShadow}, -6px -6px 12px ${barColors.lightShadow}`,
  transition: 'all 0.25s ease',
  cursor: 'pointer',

  '&:hover': {
    transform: 'translateY(-2px)',
    boxShadow: `8px 8px 16px ${barColors.darkShadow}, -8px -8px 16px ${barColors.lightShadow}`,
  },
  '&:active': {
    transform: 'translateY(2px)',
    boxShadow: `inset 6px 6px 12px ${barColors.darkShadow},
                inset -6px -6px 12px ${barColors.lightShadow}`,
  },
};

export const searchBox: SxProps = {
  flex: 1,
  maxWidth: 400,
  height: 45,
  margin: '0 2rem',
  borderRadius: '25px',
  background: barColors.background,
  boxShadow: `inset 6px 6px 12px ${barColors.darkShadow},
              inset -6px -6px 12px ${barColors.lightShadow}`,
  display: 'flex',
  alignItems: 'center',
  padding: '0 1rem',
  transition: 'all 0.25s ease',

  '&:focus-within': {
    boxShadow: `inset 3px 3px 6px ${barColors.darkShadow},
                inset -3px -3px 6px ${barColors.lightShadow}`,
  },
};

export const searchInput: React.CSSProperties = {
  width: '100%',
  border: 'none',
  outline: 'none',
  background: 'transparent',
  fontSize: '1rem',
  color: barColors.text,
};

export const logoStyle: React.CSSProperties = {
  width: 40,
  height: 40,
  filter: `
    drop-shadow(3px 3px 6px rgba(163, 177, 198, 0.6))
    drop-shadow(-3px -3px 6px rgba(255, 255, 255, 0.5))
  `,
  transition: 'all 0.25s ease',
};

export const logoHover: React.CSSProperties = {
  transform: 'translateY(-2px)',
  filter: `
    drop-shadow(4px 4px 8px rgba(163, 177, 198, 0.8))
    drop-shadow(-4px -4px 8px rgba(255, 255, 255, 0.5))
  `,
};