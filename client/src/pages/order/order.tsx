import { useLoaderData } from 'react-router-dom';
import type { Order } from '../../shared/api/order.types';
import { Card, Flex, Typography } from 'antd';

const { Text } = Typography;

export const OrderPage = () => {
  const data = useLoaderData<Order | undefined>();

  if (!data) {
    return <div>The order was not found</div>;
  }

  return (
    <Flex vertical>
      <Card>
        <Flex vertical>
          <Text>{`Order Uid: ${data.order_uid}`}</Text>
          <Text>{`Track Number: ${data.track_number}`}</Text>
          <Text>{`Entry: ${data.entry}`}</Text>
          <Text>{`Locale: ${data.locale}`}</Text>
          <Text>{`Internal Signature: ${data.internal_signature}`}</Text>
          <Text>{`Customer ID: ${data.customer_id}`}</Text>
          <Text>{`Delivery Service: ${data.delivery_service}`}</Text>
          <Text>{`Shardkey: ${data.shardkey}`}</Text>
          <Text>{`SM ID: ${data.order_uid}`}</Text>
          <Text>{`Date Created: ${data.order_uid}`}</Text>
          <Text>{`Oof Shard: ${data.order_uid}`}</Text>
        </Flex>
      </Card>
      <Flex>
        <Card>
          <Flex vertical>
            <Text>{`Name: ${data.delivery.name}`}</Text>
            <Text>{`Phone: ${data.delivery.phone}`}</Text>
            <Text>{`Zip: ${data.delivery.zip}`}</Text>
            <Text>{`City: ${data.delivery.city}`}</Text>
            <Text>{`Address: ${data.delivery.address}`}</Text>
            <Text>{`Region: ${data.delivery.region}`}</Text>
            <Text>{`Email: ${data.delivery.email}`}</Text>
          </Flex>
        </Card>
        <Card>
          <Flex vertical>
            <Text>{`Transaction: ${data.payment.transaction}`}</Text>
            <Text>{`Request ID: ${data.payment.request_id}`}</Text>
            <Text>{`Currency: ${data.payment.currency}`}</Text>
            <Text>{`Provider: ${data.payment.provider}`}</Text>
            <Text>{`Amount: ${data.payment.amount}`}</Text>
            <Text>{`Payment Dt: ${data.payment.payment_dt}`}</Text>
            <Text>{`Bank: ${data.payment.bank}`}</Text>
            <Text>{`Delivery Cost: ${data.payment.delivery_cost}`}</Text>
            <Text>{`Goods Total: ${data.payment.goods_total}`}</Text>
            <Text>{`Custom Fee: ${data.payment.transaction}`}</Text>
          </Flex>
        </Card>
      </Flex>
      <Flex>
        {data.items.map((item) => (
          <Card key={item.rid}>
            <Flex vertical>
              <Text>{`Chrt ID: ${item.chrt_id}`}</Text>
              <Text>{`Track Number: ${item.track_number}`}</Text>
              <Text>{`Price: ${item.price}`}</Text>
              <Text>{`RID: ${item.rid}`}</Text>
              <Text>{`Name: ${item.name}`}</Text>
              <Text>{`Sale: ${item.sale}`}</Text>
              <Text>{`Size: ${item.size}`}</Text>
              <Text>{`Total Price: ${item.total_price}`}</Text>
              <Text>{`NM ID: ${item.nm_id}`}</Text>
              <Text>{`Brand: ${item.brand}`}</Text>
              <Text>{`Status: ${item.status}`}</Text>
            </Flex>
          </Card>
        ))}
      </Flex>
    </Flex>
  );
};
