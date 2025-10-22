import Box from '@mui/joy/Box';
import Card from '@mui/joy/Card';
import CardCover from '@mui/joy/CardCover';
import CardContent from '@mui/joy/CardContent';
import Typography from '@mui/joy/Typography';
import ReactPlayer from 'react-player'

const VideoCard = () => {

    return (
        <Box
      component="ul"
      sx={{   }}
    >
      <Card component="li" sx={{ maxWidth: 300, minHeight:180, flexGrow: 1, "--Card-radius": "0px" }}>
        <CardCover>
          <video
            autoPlay
            loop
            muted
            poster="https://assets.codepen.io/6093409/river.jpg"
          >
            <source
              src="http://localhost:8081/video"
              type="video/mp4"
            />
          </video>
        </CardCover>
        

      </Card>
      <Card sx={{ minWidth: 300, flexGrow: 1, "--Card-radius": "0px"}}>
        <CardContent>
          <Typography
            level="body-lg"
            textColor="#fff"
            sx={{ fontWeight: 'lg', mt: { xs: 1, sm: 1 } }}
          >
            Video Information
          </Typography>
        </CardContent>
      </Card>
      <ReactPlayer src='https://www.youtube.com/watch?v=LXb3EKWsInQ' />
    </Box>
    );
}

export default VideoCard;