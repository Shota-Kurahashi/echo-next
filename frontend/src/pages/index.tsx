import { Cog8ToothIcon } from "@heroicons/react/24/outline";
import { Inter } from "next/font/google";
import { ActionIcon } from "src/components/elements/actionicon";
import { Button } from "src/components/elements/button";
import { Input } from "src/components/elements/input";
import { Loader } from "src/components/elements/loader";

const _ = Inter({ subsets: ["latin"] });

const Home = () => {
  return (
    <div>
      <Button>button</Button>
      <Loader theme="danger" />
      <Input sr="input" />
      <ActionIcon sr="setting">
        <Cog8ToothIcon className="h-6 w-6" />
      </ActionIcon>
    </div>
  );
};

export default Home;
