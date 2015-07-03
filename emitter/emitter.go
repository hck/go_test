package main

import (
  "fmt"
  "kafka_producer"
)

func main() {
	fmt.Println("starting sync producer...")

  producer := kafka_producer.Producer{}
  producer.Init([]string{"localhost:9092"})

	defer func() {
		producer.Close()
	}()

  message := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer pharetra varius turpis eget semper. Sed eleifend libero augue, sed mattis nibh tincidunt at. Curabitur sed elit mauris. Cras convallis venenatis dui in cursus. Maecenas eget mauris sed est lobortis interdum. Nullam ut est rhoncus, faucibus massa sit amet, tristique velit. Fusce et dolor nibh. Maecenas dignissim at mi non iaculis. Nullam blandit, nisi in tincidunt imperdiet, ex nibh elementum nisl, vitae ultrices augue neque vitae sapien. Fusce dapibus justo ac augue rhoncus, vitae condimentum turpis varius. Nulla facilisi. Sed eu tincidunt leo. Nam placerat velit sit amet leo lobortis vulputate. Phasellus vestibulum tempor purus ut commodo\nFusce tincidunt quam venenatis ligula molestie, non dictum nisi venenatis. Donec condimentum, enim at vulputate consectetur, nibh leo placerat risus, vitae scelerisque diam magna in nibh. Curabitur vitae sem quis magna sagittis laoreet. Ut et luctus diam. Nulla fermentum sit amet nisi vitae vestibulum. Nunc porta vel erat elementum commodo. Phasellus malesuada sapien vitae erat maximus, non placerat nisl tempus. Quisque a lectus nec orci gravida porta. Ut consequat sit amet mi sit amet posuere. Donec ac egestas lectus. Duis orci odio, mollis vel mauris in, cursus cursus odio."

  for i := 0; i < 1000000; i++ {
    producer.SendMessage("my_topic", message)
  }
}
