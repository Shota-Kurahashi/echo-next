import type { AppProps } from "next/app";
// eslint-disable-next-line import/no-unresolved
import "@/styles/tailwind.css";

const App = ({ Component, pageProps }: AppProps) => {
  return <Component {...pageProps} />;
};

export default App;
