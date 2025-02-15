using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using mini_hw1;

namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Wolf))]
public class WolfTest
{

    [TestMethod]
    public void TestIsHealthyTrue_Young()
    {
        var wolf = new Wolf("Wolf",  5, 3,8, 1, 20);
        var result = wolf.IsHealthy();
        Assert.IsTrue(result);
        Assert.IsTrue(wolf.Name == "Wolf");
    }
    [TestMethod]
    public void TestIsHealthyTrue_HasMeat()
    {
        var wolf = new Wolf("Wolf",  5, 20,8, 1, 20, 10);
        var result = wolf.IsHealthy();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestIsHealthyFalse()
    {
        var wolf = new Wolf("Wolf",  5, 20, 8, 1, 20);
        var result = wolf.IsHealthy();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_true()
    {
        var wolf = new Wolf("Wolf",  5, 3, 1, 19, 20, 5);
        var result = wolf.CanBeInContactZoo();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_False()
    {
        var wolf = new Wolf("Wolf",  5, 3, 10, 19, 20, 0);
        var result = wolf.CanBeInContactZoo();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestGiveMeat()
    {
        var wolf = new Wolf("Wolf",  5, 3, 10, 19, 20, 0);
        int meat = wolf.Meat;
        wolf.GiveFood();
        Assert.AreEqual(meat + 1, wolf.Meat);
    }
}