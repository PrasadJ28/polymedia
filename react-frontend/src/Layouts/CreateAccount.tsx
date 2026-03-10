import { TextField, Box, Button } from "@mui/material";
import { useState } from "react";

const CreateAccount = () => {
  const[username, setUsername] = useState("");
  const[email, setEmail] = useState("");
  const[firstname, setFirstname] = useState("");
  const[lastname, setLastname] = useState("");
  const[password, setPassword] = useState("");
  const[cPassword, setCPassword] = useState("");
  const createAccount = () => {
    console.log(username,email,firstname,lastname,password,cPassword);
  }
  return(
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
        label="username"
        variant="outlined"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <TextField
        label="email"
        variant="outlined"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <TextField
        label="firstname"
        variant="outlined"
        value={firstname}
        onChange={(e) => setFirstname(e.target.value)}
      />
      <TextField
        label="lastname"
        variant="outlined"
        value={lastname}
        onChange={(e) => setLastname(e.target.value)}
      />

      <TextField
        label="password"
        type="password"
        variant="outlined"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <TextField
        label="confirm password"
        type="cPassword"
        variant="outlined"
        value={cPassword}
        onChange={(e) => setCPassword(e.target.value)}
      />

      <Button variant="contained" onClick={createAccount}>
        Create
      </Button>
    </Box>
  );
}

export default CreateAccount;
