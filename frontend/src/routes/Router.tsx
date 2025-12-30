import { Routes, Route } from "react-router";
import Homepage from "../pages/Homepage";
import Notfound from "@/Notfound";
import About from "@/pages/About";
import Predictions from "@/pages/Predictions";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Homepage />} />
      <Route path="/" element={<Predictions />} />
      <Route path="/about" element={<About />} />
      <Route path="*" element={<Notfound />} />
    </Routes>
  );
};

export default AppRoutes;
