import { TextField, Box, Button } from "@mui/material";
import { loginUser } from "../Reducers/slices/loginSlice";
import { useDispatch } from "react-redux";
import { useState } from "react";
import { AppDispatch } from "../Reducers/store";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const dispatch = useDispatch<AppDispatch>();
  function login() {
    dispatch(loginUser({ username, password }));
  }
  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        gap: 2,             // Adds space between inputs
        maxWidth: '300px',  // Keeps it from getting too wide
        margin: '100px auto', // Centers it on screen
        padding: 2,
      }}
    >
      <TextField
        label="Username"
        variant="outlined"
        value={username}
        onChange={(e) => setUsername(e.target.value)} // Capture input
      />

      <TextField
        label="Password"
        type="password"
        variant="outlined"
        value={password}
        onChange={(e) => setPassword(e.target.value)} // Capture input
      />

      <Button variant="contained" onClick={login}>
        Login
      </Button>
    </Box>
  );
}

export default Login;
