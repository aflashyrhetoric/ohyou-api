select  td.*, PurchaserId = utp.UserId, BeneficiaryId = utb.UserId
into    #main
from    ExpenseDetail td
join    UserExpense utp  ON  utp.ExpenseId = td.ExpenseId
                            AND  utp.UserType = 'PURCHASER'
join    UserExpense utb  ON  utb.ExpenseId = td.ExpenseId
                            AND  utb.UserType = 'BENEFICIARY'

select  *
from    #main m
join    User u  ON u.UserID = 