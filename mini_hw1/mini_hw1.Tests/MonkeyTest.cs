using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using mini_hw1;

namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Monkey))]
public class MonkeyTest
{

    [TestMethod]
    public void TestIsHealthyTrue_Young()
    {
        var monkey = new Monkey("Monkey",  5, 3,8, 1, 20, 0);
        var result = monkey.IsHealthy();
        Assert.IsTrue(result);
        Assert.IsTrue(monkey.Name == "Monkey");
    }
    [TestMethod]
    public void TestIsHealthyTrue_HasBananas()
    {
        var monkey = new Monkey("Monkey",  5, 20,8, 1, 20, 5);
        var result = monkey.IsHealthy();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestIsHealthyFalse()
    {
        var monkey = new Monkey("Monkey",  5, 20, 8, 1, 20, 0);
        var result = monkey.IsHealthy();
        Assert.IsFalse(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_true()
    {
        var monkey = new Monkey("Monkey",  5, 3, 9, 19, 20, 5);
        var result = monkey.CanBeInContactZoo();
        Assert.IsTrue(result);
    }
    [TestMethod]
    public void TestCanBeInContactZoo_False()
    {
        var monkey = new Monkey("Rabbit",  5, 3, 1, 19, 20, 0);
        var result = monkey.CanBeInContactZoo();
        Assert.IsFalse(result);
    }

    [TestMethod]
    public void TestGiveBananas()
    {
        var monkey = new Monkey("Monkey",  5, 3, 9, 19, 20, 5);
        int bananas = monkey.Bananas;
        monkey.GiveFood();
        Assert.AreEqual(bananas + 1, monkey.Bananas);
    }
}