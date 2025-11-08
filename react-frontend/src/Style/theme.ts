import { extendTheme } from '@mui/joy/styles';

// ---- STEP 1: Module Augmentation ----
declare module '@mui/joy/styles' {
  interface Theme {
    custom: {
      shadows: {
        raised: string;
        inset: string;
        flat: string;
      };
      surfaces: {
        background: string;
        surface: string;
      };
      radius: {
        sm: string;
        md: string;
        lg: string;
      };
    };
  }

  interface ThemeScales {
    custom: Theme['custom'];
  }

  interface ThemeOptions {
    custom?: {
      shadows?: Partial<Theme['custom']['shadows']>;
      surfaces?: Partial<Theme['custom']['surfaces']>;
      radius?: Partial<Theme['custom']['radius']>;
    };
  }
}

// ---- STEP 2: Define Pastel Neumorphic Theme ----
export const theme = extendTheme({
  colorSchemes: {
    light: {
      palette: {
        primary: { solidBg: '#A8C5E8', solidHoverBg: '#91B3DC' },
        neutral: { solidBg: '#E6E9EF', solidHoverBg: '#DADFE7' },
        text: { primary: '#2D3A4A', secondary: '#505C6B' },
        background: { body: '#E6E9EF' },
      },
    },
    dark: {
      palette: {
        primary: { solidBg: '#91B3DC', solidHoverBg: '#7FA1CC' },
        neutral: { solidBg: '#2B2E36', solidHoverBg: '#373B43' },
        text: { primary: '#E9EEF5', secondary: '#C2C7CF' },
        background: { body: '#2B2E36' },
      },
    },
  },

  fontFamily: {
    body: "'Inter', sans-serif",
  },

  // ---- STEP 3: Custom Tokens ----
  custom: {
    shadows: {
      flat: 'none',
      // 🩵 Softer dual shadows for raised look
      raised: '6px 6px 12px #C7CCD6, -6px -6px 12px #FAFBFF',
      inset: 'inset 6px 6px 12px #C7CCD6, inset -6px -6px 12px #FAFBFF',
    },
    surfaces: {
      background: '#E6E9EF', // base background
      surface: '#E6E9EF', // card surface
    },
    radius: {
      sm: '10px',
      md: '16px',
      lg: '24px',
    },
  },

  // ---- STEP 4: Component Overrides ----
  components: {
    JoyCard: {
      styleOverrides: {
        root: ({ theme }) => ({
          backgroundColor: theme.custom.surfaces.surface,
          borderRadius: theme.custom.radius.lg,
          boxShadow: theme.custom.shadows.raised,
          transition: 'all 0.25s ease',
          '&:hover': {
            boxShadow: '8px 8px 16px #C3C8D2, -8px -8px 16px #FFFFFF',
            transform: 'translateY(-3px)',
          },
          '&:active': {
            boxShadow: theme.custom.shadows.inset,
            transform: 'translateY(2px)',
          },
        }),
      },
    },
    JoyButton: {
      styleOverrides: {
        root: ({ theme }) => ({
          backgroundColor: theme.custom.surfaces.surface,
          borderRadius: theme.custom.radius.md,
          boxShadow: theme.custom.shadows.raised,
          transition: 'all 0.25s ease',
          color: theme.colorSchemes.light.palette.text.primary,
          '&:hover': {
            boxShadow: '8px 8px 16px #C3C8D2, -8px -8px 16px #FFFFFF',
            transform: 'scale(1.03)',
          },
          '&:active': {
            boxShadow: theme.custom.shadows.inset,
            transform: 'scale(0.97)',
          },
        }),
      },
    },
    JoyInput: {
      styleOverrides: {
        root: ({ theme }) => ({
          backgroundColor: theme.custom.surfaces.surface,
          borderRadius: theme.custom.radius.sm,
          boxShadow: theme.custom.shadows.inset,
          border: 'none',
          '&:focus-within': {
            boxShadow: 'inset 3px 3px 8px #C7CCD6, inset -3px -3px 8px #FFFFFF',
          },
        }),
      },
    },
  },
});
