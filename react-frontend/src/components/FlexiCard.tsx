import { useEffect, useState } from "react";
import { Card, Typography } from "@mui/joy";
import { cardStyle, textContainer, textClamp } from "../Style/components/FlexiCard.styles";

export interface ContentType {
  type: "text" | "image" | "video";
  text?: string;
  src?: string;
  caption?: string;
}

export default function FlexiCard({ content }: { content: ContentType }) {
  const [orientation, setOrientation] = useState<"horizontal" | "vertical">("vertical");

  useEffect(() => {
    if (content.type === "image" && content.src) {
      const img = new Image();
      img.src = content.src;
      img.onload = () => setOrientation(img.width > img.height ? "horizontal" : "vertical");
    }
    if (content.type === "video" && content.src) {
      const video = document.createElement("video");
      video.preload = "metadata";
      video.src = content.src;
      video.onloadedmetadata = () =>
        setOrientation(video.videoWidth > video.videoHeight ? "horizontal" : "vertical");
    }
  }, [content]);

  return (
    <Card
      sx={{
        ...cardStyle(orientation, "auto"),
        width: "100%",
        maxWidth: 220, // 👈 match grid column width
        overflow: "hidden",
        padding: 0,
      }}
    >
      {content.type === "image" && (
        <img
          src={content.src}
          alt="media"
          style={{
            width: "100%",
            height: "auto",
            maxHeight: 200, // 👈 keeps cards compact
            objectFit: "cover",
            display: "block",
          }}
        />
      )}

      {content.type === "video" && (
        <video
          src={content.src}
          controls
          style={{
            width: "100%",
            height: "auto",
            maxHeight: 200, // 👈 limits video size
            objectFit: "cover",
            display: "block",
          }}
        />
      )}

      {content.type === "text" && (
        <div style={textContainer}>
          <Typography level="body-md" sx={textClamp}>
            {content.text}
          </Typography>
        </div>
      )}

      {content.caption && (
        <Typography level="body-sm" sx={{ padding: "6px 8px" }}>
          {content.caption}
        </Typography>
      )}
    </Card>
  );
}
