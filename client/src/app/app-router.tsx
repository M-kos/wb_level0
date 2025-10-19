import { createBrowserRouter } from 'react-router';
import { Home } from '../pages/home/home.tsx';
import { routes } from '../shared/routes/routes.ts';
import { OrderPage } from '../pages/order/order.tsx';
import { RouterProvider } from 'react-router-dom';
import { orderRepository } from '../shared/api/order.repository.ts';

const appRouter = createBrowserRouter([
  { path: routes.home, Component: Home },
  {
    path: routes.order,
    Component: OrderPage,
    loader: async ({ params }) => {
      if (!params.orderId) {
        return null;
      }

      const order = await orderRepository.getOrder(params.orderId);

      return {
        order,
      };
    },
  },
]);

export const AppRouter = () => {
  return <RouterProvider router={appRouter} />;
};
