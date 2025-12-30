import { BrowserRouter } from "react-router";
import MainLayout from "./Layout";
import AppRoutes from "./routes/Router";

const App = () => {
  return (
    <BrowserRouter>
      <MainLayout>
        <AppRoutes />
      </MainLayout>
    </BrowserRouter>
  );
};

export default App;
