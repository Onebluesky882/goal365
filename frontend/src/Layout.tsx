import { NavLink } from "./components/Footer";
import Header from "./components/Header";

const MainLayout = ({ children }: React.PropsWithChildren) => {
  return (
    <>
      <Header />
      <main className="container">{children}</main>
      <NavLink to="/">Home</NavLink>
    </>
  );
};

export default MainLayout;
