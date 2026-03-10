declare module '*.css';
declare module '*.scss';
declare module '*.sass';
declare module '*.less';

interface ImportMetaEnv {
  readonly VITE_API_URL: string;
  readonly VITE_API_KEY: string;
  // more env variables...
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
