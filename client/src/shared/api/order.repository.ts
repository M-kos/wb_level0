import type { Order } from './order.types.ts';

export const orderRepository = {
  getOrder: async (orderId: string): Promise<Order | undefined> => {
    try {
      const response = await fetch(`/api/orders/${orderId}`, {});

      return await response.json();
    } catch (error) {
      console.error(error);
    }
  },
};
