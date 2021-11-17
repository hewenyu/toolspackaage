package kafka_plus

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

// rebalancedCallback is called on each group rebalance to assign additional
// partitions, or remove existing partitions, from the consumer's current
// assignment.
//
// The application may use this optional callback to inspect the assignment,
// alter the initial start offset (the .Offset field of each assigned partition),
// and read/write offsets to commit to an alternative store outside of Kafka.
func rebalancedCallback(c *kafka.Consumer, event kafka.Event) error {
	
	switch ev := event.(type) {
	case kafka.AssignedPartitions:
		_, err := fmt.Fprintf(os.Stderr,
			"%% %s rebalance: %d new partition(s) assigned: %v\n",
			c.GetRebalanceProtocol(), len(ev.Partitions),
			ev.Partitions)
		if err != nil {
			return err
		}
		
		// The application may update the start .Offset of each
		// assigned partition and then call IncrementalAssign().
		// Even though this example does not alter the offsets we
		// provide the call to IncrementalAssign() as an example.
		err = c.IncrementalAssign(ev.Partitions)
		if err != nil {
			panic(err)
		}
	
	case kafka.RevokedPartitions:
		_, err := fmt.Fprintf(os.Stderr,
			"%% %s rebalance: %d partition(s) revoked: %v\n",
			c.GetRebalanceProtocol(), len(ev.Partitions),
			ev.Partitions)
		if err != nil {
			return err
		}
		if c.AssignmentLost() {
			// Our consumer has been kicked out of the group and the
			// entire assignment is thus lost.
			_, err := fmt.Fprintf(os.Stderr, "%% Current assignment lost!\n")
			if err != nil {
				return err
			}
		}
		
		// The client automatically calls IncrementalUnassign() unless
		// the callback has already called that method.
	}
	
	return nil
}