FROM rabbitmq:3.8-management

# Enabled rabbitmq plugins
# * rabbitmq_management
# * rabbitmq_management_agent
# * rabbitmq_mqtt
# * rabbitmq_stomp
# * rabbitmq_delayed_message_exchange
# * rabbitmq_web_dispatch
# * rabbitmq_amqp1_0

# install curl
RUN apt-get update && apt-get install -y curl && rm -rf /var/cache/apt/lists

# copy the delayed_message_exchange plugin
RUN curl -LO https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.8.17/rabbitmq_delayed_message_exchange-3.8.17.8f537ac.ez
RUN mv ./rabbitmq_delayed_message_exchange-3.8.17.8f537ac.ez /opt/rabbitmq/plugins/

# install plugins
RUN rabbitmq-plugins enable rabbitmq_management rabbitmq_management_agent rabbitmq_mqtt rabbitmq_stomp rabbitmq_delayed_message_exchange rabbitmq_web_dispatch rabbitmq_amqp1_0
