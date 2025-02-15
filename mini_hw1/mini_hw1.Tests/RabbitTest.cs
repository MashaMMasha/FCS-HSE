using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using mini_hw1;

namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Rabbit))]
public class RabbitTest
{

    [TestMethod]
    public void TestIsHealthyTrue_Young()
    {
        var rabbit = new Rabbit("Rabbit",  5, 3,8, 1, 20, 0);
        var result = rabbit.IsHealthy();
        Assert.IsTrue(result);
        Assert.IsTrue(rabbit.Name == "Rabbit");
    }
    [TestMethod]
    public void TestIsHealthyTrue_HasCarrot()
    {
        var rabbit = new Rabbit("Rabbit",  5, 20,8, 1, 20, 5);
        var result = rabbit.IsHealthy();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestIsHealthyFalse()
    {
        var rabbit = new Rabbit("Rabbit",  5, 20, 8, 1, 20, 0);
        var result = rabbit.IsHealthy();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_true()
    {
        var rabbit = new Rabbit("Rabbit",  5, 3, 8, 19, 20, 5);
        var result = rabbit.CanBeInContactZoo();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_False()
    {
        var rabbit = new Rabbit("Rabbit",  5, 3, 1, 19, 20, 0);
        var result = rabbit.CanBeInContactZoo();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestGiveCarrot()
    {
        var rabbit = new Rabbit("Rabbit",  5, 3, 9, 19, 20, 5);
        int carrot = rabbit.Carrot;
        rabbit.GiveFood();
        Assert.AreEqual(carrot + 1, rabbit.Carrot);
    }
}