import type { Order } from './order.types.ts';

const host = 'localhost:8083';

export const orderRepository = {
  getOrder: async (orderId: string): Promise<Order | undefined> => {
    try {
      const response = await fetch(`${host}/api/orders/${orderId}`, {});
      return await response.json();
    } catch (error) {
      console.error(error);
    }
  },
};
