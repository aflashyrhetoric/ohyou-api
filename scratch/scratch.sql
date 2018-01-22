select  td.*, PurchaserId = utp.UserId, BeneficiaryId = utb.UserId
into    #main
from    TransactionDetail td
join    UserTransaction utp  ON  utp.TransactionId = td.TransactionId
                            AND  utp.UserType = 'PURCHASER'
join    UserTransaction utb  ON  utb.TransactionId = td.TransactionId
                            AND  utb.UserType = 'BENEFICIARY'

select  *
from    #main m
join    User u  ON u.UserID = 