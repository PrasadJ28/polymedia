import { useEffect, useRef } from "react";
import FlexiCard, { ContentType } from "../components/FlexiCard";

interface MasonryLayoutProps {
  items: ContentType[];
}

export default function MasonryLayout({ items }: MasonryLayoutProps) {
  const gridRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const resizeAllGridItems = () => {
      const grid = gridRef.current;
      if (!grid) return;

      const rowHeight = parseInt(
        window.getComputedStyle(grid).getPropertyValue("grid-auto-rows")
      );
      const rowGap = parseInt(
        window.getComputedStyle(grid).getPropertyValue("grid-row-gap")
      );

      grid.querySelectorAll(".masonry-item").forEach((item) => {
        const el = item as HTMLElement;
        const rowSpan = Math.ceil(
          (el.scrollHeight + rowGap) / (rowHeight + rowGap)
        );
        el.style.gridRowEnd = `span ${rowSpan}`;
      });
    };

    const handleMediaLoad = () => resizeAllGridItems();

    window.addEventListener("resize", resizeAllGridItems);

    // Re-measure whenever an image or video loads
    const grid = gridRef.current;
    if (grid) {
      grid.querySelectorAll("img, video").forEach((el) =>
        el.addEventListener("load", handleMediaLoad)
      );
    }

    resizeAllGridItems();

    return () => {
      window.removeEventListener("resize", resizeAllGridItems);
      if (grid) {
        grid.querySelectorAll("img, video").forEach((el) =>
          el.removeEventListener("load", handleMediaLoad)
        );
      }
    };
  }, [items]);

  return (
    <div
      ref={gridRef}
      style={{
        display: "grid",
        gridTemplateColumns: "repeat(auto-fill, minmax(220px, 1fr))",
        gridAutoRows: "10px",
        gap: "15px",
        justifyContent: "center",
        padding: "20px",
        width: "90vw",
        maxWidth: "1200px",
        margin: "0 auto",
        overflowY: "auto",
      }}
    >
      {items.map((item, index) => (
        <div className="masonry-item" key={index}>
          <FlexiCard content={item} />
        </div>
      ))}
    </div>
  );
}
