namespace BankSystem;

public class BankAccount
{
    public readonly Guid id;
    public readonly string name;
    public double balance;

    public BankAccount(string name)
    {
        id = Guid.NewGuid();
        this.name = name;
        balance = 0;
    }
    public BankAccount(string name, double balance)
    {
        id = Guid.NewGuid();
        this.name = name;
        this.balance = balance;
    }
}