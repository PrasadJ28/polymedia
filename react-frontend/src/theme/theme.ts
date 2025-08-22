import { createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#FF9F00", // your new primary color (Tailwind blue-800 as example)
      contrastText: "#ffffff", // text color on primary
    },
    secondary: {
      main: "#f59e0b", // optional secondary (Tailwind amber-500)
    },
  },
});

export default theme;
