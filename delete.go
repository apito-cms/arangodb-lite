package arangodb

import (
	"fmt"
	//"github.com/apex/log"
	//"github.com/thedanielforum/arangodb/errc"
)

func (c *Connection) Delete(collectionName string ,docHandle string) error {

	endPoint := fmt.Sprintf("/_db/%s/_api/document/%s/%s", c.db, collectionName,docHandle)

	_,err := c.deleteReq(endPoint)
	if err != nil {
		//log.WithError(err).Info(arangodb.ErrorCodeInvalidEdgeAttribute.Error().Error())
		//return arangodb.ErrorCodeDocNotExist.Error()
	}
	return nil
}