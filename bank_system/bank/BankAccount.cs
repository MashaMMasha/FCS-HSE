namespace BankSystem;

public class BankAccount
{
    readonly Guid id;
    readonly string name;
    double balance;

    BankAccount(string name)
    {
        id = Guid.NewGuid();
        this.name = name;
        balance = 0;
    }
    BankAccount(string name, double balance)
    {
        id = Guid.NewGuid();
        this.name = name;
        this.balance = balance;
    }
}