import XCTest
import SwiftTreeSitter
import TreeSitterAwk

final class TreeSitterAwkTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_awk())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Awk grammar")
    }
}
