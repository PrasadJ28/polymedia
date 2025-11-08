
import { Box } from "@mui/joy";
import FlexiCard from "../components/FlexiCard";
import type { ContentType } from '../components/FlexiCard';
import StudioBar from "../components/StudioBar";

const Dashboard = () => {
 const cards: ContentType[] = [
    { type: 'text', text: 'This is a long paragraph that will truncate.' },
    { type: 'image', src: 'https://picsum.photos/seed/picsum/200/300', caption: 'Sample image' },
    { type: 'video', src: 'https://www.w3schools.com/html/mov_bbb.mp4', caption: 'Sample video' },
  ];

    return (
        <div>
            <StudioBar />
        
        <Box
      sx={{
        backgroundColor: '#bababaff', // 👈 pastel background
        minHeight: '100vh',
        display: 'flex',
        flexWrap: 'wrap',
        gap: 3,
        justifyContent: 'center',
        alignItems: 'center',
        padding: 4,
      }}
    >
        
    <div style={{ display: 'flex', gap: 16, flexWrap: 'wrap' }}>
      {cards.map((item, i) => (
        <FlexiCard key={i} content={item} />
      ))}
    </div>
         </Box>
         </div>
    );
}

export default Dashboard;