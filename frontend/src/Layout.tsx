import { NavLink } from "./components/Footer";
import Header from "./components/Header";

const MainLayout = ({ children }: React.PropsWithChildren) => {
  return (
    <div className="">
      <Header />
      <main className="">{children}</main>
      <NavLink to="/">Home</NavLink>
    </div>
  );
};

export default MainLayout;
