import { isRouteErrorResponse, useRouteError } from "react-router-dom";

export default function ErrorLayout() {
  const err = useRouteError();
  console.error("Route error:", err);

  if (isRouteErrorResponse(err)) {
    return (
      <div className="p-6">
        <h1>{err.status} {err.statusText}</h1>
        {err.data ? <pre>{JSON.stringify(err.data, null, 2)}</pre> : null}
      </div>
    );
  }

  return (
    <div className="p-6">
      <h1>Something went wrong</h1>
      <pre>{err instanceof Error ? err.message : String(err)}</pre>
    </div>
  );
}