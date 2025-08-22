import * as React from "react";
import {
  AppBar, Box, Toolbar, IconButton, Typography, Menu, Avatar,
  Tooltip, MenuItem, SvgIcon, TextField
} from "@mui/material";
import { ReactComponent as Wolf } from "../assets/wolf.svg";
import "../theme/styles/navbar.css"; // 👈 import the Tailwind-based stylesheet

const settings = ["Profile", "Account", "Dashboard", "Logout"];

function NavBar() {
  const [anchorElUser, setAnchorElUser] = React.useState<null | HTMLElement>(null);

  const handleOpenUserMenu = (e: React.MouseEvent<HTMLElement>) => setAnchorElUser(e.currentTarget);
  const handleCloseUserMenu = () => setAnchorElUser(null);

  return (
    <AppBar position="static" color="primary">
      <Toolbar disableGutters className="nav-toolbar">
        {/* LEFT: Logo + Brand */}
        <Box className="logo-group">
          <SvgIcon component={Wolf} inheritViewBox className="logo-icon"  />
          <Typography component="a" href="/"  className="brand text-6xl font-bold leading-tight text-white no-underline">
            howl
          </Typography>
        </Box>

        {/* CENTER: Search */}
        <Box className="center">
          <TextField
            id="search"
            placeholder="Search..."
            variant="outlined"
            size="small"
            className="search"
          />
        </Box>

        {/* RIGHT: Avatar menu */}
        <Box className="right-group">
          <Tooltip title="Open settings">
            <IconButton onClick={handleOpenUserMenu} className="p-0">
              <Avatar alt="Remy Sharp" src="/static/images/avatar/2.jpg" />
            </IconButton>
          </Tooltip>
          <Menu
            className="mt-11"
            id="menu-appbar"
            anchorEl={anchorElUser}
            anchorOrigin={{ vertical: "top", horizontal: "right" }}
            keepMounted
            transformOrigin={{ vertical: "top", horizontal: "right" }}
            open={Boolean(anchorElUser)}
            onClose={handleCloseUserMenu}
          >
            {settings.map((setting) => (
              <MenuItem key={setting} onClick={handleCloseUserMenu}>
                <Typography className="text-center">{setting}</Typography>
              </MenuItem>
            ))}
          </Menu>
        </Box>
      </Toolbar>
    </AppBar>
  );
}

export default NavBar;
