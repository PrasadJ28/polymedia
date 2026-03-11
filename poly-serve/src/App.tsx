import { useRef, useState } from "react";
import "./index.css";


export function App() {
  const inputFile = useRef<HTMLInputElement>(null);
  const [status, setStatus]     = useState<string>("");
  const [progress, setProgress] = useState<number>(0);
  const [error, setError]       = useState<string>("");

  const PART_SIZE = 10 * 1024 * 1024;
  const startUpload = async (filename: string, filesize: number) => {
  const response = await fetch("http://localhost:8080/upload/start", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ filename, filesize }),
  });

  if (!response.ok) {
    throw new Error("Failed to start upload");
  }

  return response.json() as Promise<{
    uploadId: string;
    totalParts: number;
    partSize: number;
  }>;
};

const presignAndUploadPart = async (
  uploadId: string,
  key: string,
  partNumber: number,
  chunk: Blob
): Promise<{ partNumber: number; etag: string }> => {

  // Step 1 — get the presigned URL from your Go server
  const presignResponse = await fetch(
    `http://localhost:8080/upload/presign?uploadId=${uploadId}&key=${key}&partNumber=${partNumber}`
  );

  if (!presignResponse.ok) {
    throw new Error(`Failed to get presigned URL for part ${partNumber}`);
  }

  const { presignedURL } = await presignResponse.json();

  // Step 2 — PUT the chunk directly to Minio using the presigned URL
  const uploadResponse = await fetch(presignedURL, {
    method: "PUT",
    body: chunk,
  });

  if (!uploadResponse.ok) {
    throw new Error(`Failed to upload part ${partNumber}`);
  }

  // Step 3 — extract the ETag from Minio's response headers
  const etag = uploadResponse.headers.get("ETag") ?? "";

  return { partNumber, etag };
};
const completeUpload = async (
  uploadId: string,
  key: string,
  parts: { partNumber: number; etag: string }[]
): Promise<void> => {

  const response = await fetch("http://localhost:8080/upload/complete", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ uploadId, key, parts }),
  });

  if (!response.ok) {
    throw new Error("Failed to complete upload");
  }
};
  const uploadFile = async () => {
  // Guard — do nothing if no file selected
  if (!inputFile.current?.files?.[0]) {
    setError("Please select a file first");
    return;
  }

  const file = inputFile.current.files[0];
  setError("");
  setProgress(0);

  try {
    // Step 1 — start the multipart session
    setStatus("Starting upload...");
    const { uploadId, totalParts } = await startUpload(file.name, file.size);

    // Step 2 — upload each chunk
    const parts: { partNumber: number; etag: string }[] = [];

    for (let i = 0; i < totalParts; i++) {
      const start      = i * PART_SIZE;
      const end        = Math.min(start + PART_SIZE, file.size);
      const chunk      = file.slice(start, end);
      const partNumber = i + 1;

      setStatus(`Uploading part ${partNumber} of ${totalParts}...`);

      const part = await presignAndUploadPart(uploadId, file.name, partNumber, chunk);
      parts.push(part);

      setProgress(Math.round((partNumber / totalParts) * 100));
    }

    // Step 3 — assemble the file
    setStatus("Completing upload...");
    await completeUpload(uploadId, file.name, parts);

    setStatus("Upload complete!");
    setProgress(100);

  } catch (err) {
    setError(err instanceof Error ? err.message : "Upload failed");
    setStatus("");
  }
};
  return (
    <div className="app">
      <input type="file" id="file" ref={inputFile} />
      <button onClick={uploadFile}>Upload</button>
      {status   && <p>{status}</p>}
    {progress > 0 && <progress value={progress} max={100} />}
    {error    && <p style={{ color: "red" }}>{error}</p>}
    </div>
  );
}

export default App;
