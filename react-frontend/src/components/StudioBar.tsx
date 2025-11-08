import { useState } from 'react';
import { Box } from '@mui/joy';
import { AccountCircle, Search } from '@mui/icons-material';
import { ReactComponent as Logo } from '../assets/logo.svg';
import {
  barContainer,
  elevatedCircle,
  logoHover,
  logoStyle,
  searchBox,
  searchInput,
} from '../Style/components/StudioBar.styles';

export default function StudioBar() {
     const [isHovered, setIsHovered] = useState(false);
  return (
    <Box sx={barContainer}>
      {/* Left: Logo (elevated by default) */}
      <div
        style={{
          ...logoStyle,
          ...(isHovered ? logoHover : {}),
        }}
        onMouseEnter={() => setIsHovered(true)}
        onMouseLeave={() => setIsHovered(false)}
      >
        <Logo width={80} height={80} />
      </div>

      {/* Center: Neumorphic Search Bar */}
      <Box sx={searchBox}>
        <Search
          style={{ marginRight: '10px', color: 'rgba(0,0,0,0.4)' }}
          fontSize="small"
        />
        <input type="text" placeholder="Search..." style={searchInput} />
      </Box>

      {/* Right: Account Icon */}
      <Box sx={elevatedCircle}>
        <AccountCircle style={{ fontSize: 28, color: '#002758' }} />
      </Box>
    </Box>
  );
}
